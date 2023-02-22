package impl

import (
	"context"
	"fmt"
	"strings"

	"github.com/renmcc/go-cmdb/apps/host"
)

func (i *HostServiceImpl) CreateHost(ctx context.Context, ins *host.Host) (*host.Host, error) {
	// 默认值填充
	ins.InjectDefault()

	// 把数据入库到 resource表和host表
	// 一次需要往2个表录入数据, 我们需要2个操作 要么都成功，要么都失败, 事务的逻辑
	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("start tx error, %s", err)
	}

	// 通过Defer处理事务提交方式
	// 1. 无报错，则Commit 事务
	// 2. 有报错, 则Rollback 事务
	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				i.log.Named("CreateHost").Error("rollback error, %s", err)
			}
		} else {
			if err := tx.Commit(); err != nil {
				i.log.Named("CreateHost").Error("commit error, %s", err)
			}
		}
	}()

	// 插入Resource数据
	rstmt, err := tx.PrepareContext(ctx, InsertResourceSQL)
	if err != nil {
		return nil, err
	}
	defer rstmt.Close()

	_, err = rstmt.ExecContext(ctx,
		ins.Id, ins.Vendor, ins.Region, ins.CreateAt, ins.ExpireAt, ins.Type,
		ins.Name, ins.Description, ins.Status, ins.UpdateAt, ins.SyncAt, ins.Account, ins.PublicIP,
		ins.PrivateIP,
	)
	if err != nil {
		return nil, err
	}

	// 插入Describe 数据
	dstmt, err := tx.PrepareContext(ctx, InsertDescribeSQL)
	if err != nil {
		return nil, err
	}
	defer dstmt.Close()

	_, err = dstmt.ExecContext(ctx,
		ins.Id, ins.CPU, ins.Memory, ins.GPUAmount, ins.GPUSpec,
		ins.OSType, ins.OSName, ins.SerialNumber,
	)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *HostServiceImpl) QueryHost(ctx context.Context, req *host.QueryHostRequest) (*host.HostSet, error) {
	args := []interface{}{req.Name, req.Description, req.PrivateIp, req.PublicIp, req.OffSet(), req.GetPageSize()}

	i.log.Named("QueryHost").Infof("query sql: %s; %v", QueryHostSQL, args)

	// query stmt, 构建一个Prepare语句
	stmt, err := i.db.PrepareContext(ctx, QueryHostSQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	set := host.NewHostSet()
	for rows.Next() {
		// 没扫描一行,就需要读取出来
		ins := host.NewHost()
		if err := rows.Scan(&ins.Id, &ins.Vendor, &ins.Region, &ins.CreateAt, &ins.ExpireAt,
			&ins.Type, &ins.Name, &ins.Description, &ins.Status, &ins.UpdateAt, &ins.SyncAt,
			&ins.Account, &ins.PublicIP, &ins.PrivateIP,
			&ins.CPU, &ins.Memory, &ins.GPUSpec, &ins.GPUAmount, &ins.OSType, &ins.OSName, &ins.SerialNumber,
		); err != nil {
			return nil, err
		}
		set.Add(ins)
	}

	// total统计
	set.Total = len(set.Items)

	return set, nil
}

func (i *HostServiceImpl) DescribeHost(ctx context.Context, req *host.DescribeHostRequest) (*host.Host, error) {
	args := []interface{}{req.Id}
	i.log.Named("DescribeHost").Infof("query sql: %s; %v", DescribeHostSQL, args)

	// query stmt, 构建一个Prepare语句
	stmt, err := i.db.PrepareContext(ctx, DescribeHostSQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// 取出数据，赋值结构体
	ins := host.NewHost()
	err = stmt.QueryRowContext(ctx, args...).Scan(
		&ins.Id, &ins.Vendor, &ins.Region, &ins.CreateAt, &ins.ExpireAt,
		&ins.Type, &ins.Name, &ins.Description, &ins.Status, &ins.UpdateAt, &ins.SyncAt,
		&ins.Account, &ins.PublicIP, &ins.PrivateIP,
		&ins.CPU, &ins.Memory, &ins.GPUSpec, &ins.GPUAmount, &ins.OSType, &ins.OSName, &ins.SerialNumber,
	)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, nil
		}
		return nil, err
	}

	return ins, nil
}

func (i *HostServiceImpl) UpdateHost(ctx context.Context, ins *host.Host) (*host.Host, error) {
	// 开启一个事务
	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	// 通过Defer处理事务提交方式
	// 1. 无报错，则Commit 事务
	// 2. 有报错，则Rollback 事务
	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				i.log.Error("rollback error, %s", err.Error())
			}
		} else {
			if err := tx.Commit(); err != nil {
				i.log.Error("commit error, %s", err.Error())
			}
		}
	}()

	// 更新 Resource表
	resStmt, err := tx.PrepareContext(ctx, updateResourceSQL)
	if err != nil {
		return nil, err
	}
	_, err = resStmt.ExecContext(ctx, ins.Vendor, ins.Region, ins.ExpireAt, ins.Name, ins.Description, ins.Id)
	if err != nil {
		return nil, err
	}

	// 更新 Host表
	hostStmt, err := tx.PrepareContext(ctx, updateHostSQL)
	if err != nil {
		return nil, err
	}
	_, err = hostStmt.ExecContext(ctx, ins.CPU, ins.Memory, ins.Id)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *HostServiceImpl) DeleteHost(ctx context.Context, req *host.DeleteHostRequest) error {
	// 开启一个事务
	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// 通过Defer处理事务提交方式
	// 1. 无报错，则Commit 事务
	// 2. 有报错，则Rollback 事务
	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				i.log.Error("rollback error, %s", err.Error())
			}
		} else {
			if err := tx.Commit(); err != nil {
				i.log.Error("commit error, %s", err.Error())
			}
		}
	}()

	// 更新 Resource表
	resStmt, err := tx.PrepareContext(ctx, deleteHostSQL)
	if err != nil {
		return err
	}
	_, err = resStmt.ExecContext(ctx, req.Id)
	if err != nil {
		return err
	}

	return nil
}
