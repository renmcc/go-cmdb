package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/logger/zap"
	"github.com/renmcc/go-cmdb/apps/host"
)

const (
	InsertResourceSQL = `
	INSERT INTO resource (id,vendor,region,create_at,expire_at,type,name,description,status,update_at,sync_at,accout,public_ip,private_ip)
	VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?);
	`
	InsertDescribeSQL = `
	INSERT INTO host (resource_id, cpu, memory, gpu_amount, gpu_spec, os_type, os_name, serial_number)
	VALUES( ?,?,?,?,?,?,?,? );
	`
)

// 把Host对象保存到数据内, 数据的一致性
func (i *HostServiceImpl) save(ctx context.Context, ins *host.Host) error {
	var (
		err error
	)

	// 把数据入库到 resource表和host表
	// 一次需要往2个表录入数据, 我们需要2个操作 要么都成功，要么都失败, 事务的逻辑
	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("start tx error, %s", err)
	}

	// 通过Defer处理事务提交方式
	// 1. 无报错，则Commit 事务
	// 2. 有报错, 则Rollback 事务
	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				zap.L().Named("Host").Error("rollback error, %s", err)
			}
		} else {
			if err := tx.Commit(); err != nil {
				zap.L().Named("Host").Error("commit error, %s", err)
			}
		}
	}()

	// 插入Resource数据
	rstmt, err := tx.PrepareContext(ctx, InsertResourceSQL)
	if err != nil {
		return err
	}
	defer rstmt.Close()

	_, err = rstmt.ExecContext(ctx,
		ins.Id, ins.Vendor, ins.Region, ins.CreateAt, ins.ExpireAt, ins.Type,
		ins.Name, ins.Description, ins.Status, ins.UpdateAt, ins.SyncAt, ins.Account, ins.PublicIP,
		ins.PrivateIP,
	)
	if err != nil {
		return err
	}

	// 插入Describe 数据
	dstmt, err := tx.PrepareContext(ctx, InsertDescribeSQL)
	if err != nil {
		return err
	}
	defer dstmt.Close()

	_, err = dstmt.ExecContext(ctx,
		ins.Id, ins.CPU, ins.Memory, ins.GPUAmount, ins.GPUSpec,
		ins.OSType, ins.OSName, ins.SerialNumber,
	)
	if err != nil {
		return err
	}

	return nil
}
