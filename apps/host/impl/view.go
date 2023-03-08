package impl

import (
	"context"
	"strings"

	"github.com/renmcc/go-cmdb/apps/host"
)

func (i *HostServiceImpl) CreateHost(ctx context.Context, ins *host.Host) error {
	// 默认值填充
	ins.InsertDefault()

	// 开启事务
	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		return err
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
	rstmt, err := tx.PrepareContext(ctx, InsertHostSQL)
	if err != nil {
		return err
	}
	defer rstmt.Close()
	_, err = rstmt.ExecContext(ctx,
		ins.Xid, ins.Status, ins.CreateAt, ins.CreateBy, ins.ResourceId, ins.Vendor,
		ins.Name, ins.Region, ins.ExpireAt, ins.PublicIP, ins.PrivateIP, ins.CPU,
		ins.Memory, ins.OSType, ins.OSName, ins.SerialNumber,
	)
	if err != nil {
		return err
	}

	return nil
}

func (i *HostServiceImpl) ListHost(ctx context.Context, req *host.ListHostRequest) (*host.HostSet, error) {
	args := []interface{}{req.SerialNumber, req.PrivateIp, req.OffSet(), req.GetPageSize()}

	i.log.Named("ListHost").Infof("query sql: %s; %v", ListHostSQL, args)

	// query stmt, 构建一个Prepare语句
	stmt, err := i.db.PrepareContext(ctx, ListHostSQL)
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
		if err := rows.Scan(&ins.Xid, &ins.Status, &ins.CreateAt, &ins.CreateBy, &ins.UpdateAt, &ins.UpdateBy, &ins.DeleteAt, &ins.DeleteBy, &ins.ResourceId, &ins.Vendor,
			&ins.Name, &ins.Region, &ins.ExpireAt, &ins.PublicIP, &ins.PrivateIP, &ins.CPU, &ins.Memory, &ins.OSType, &ins.OSName, &ins.SerialNumber,
		); err != nil {
			return nil, err
		}
		set.Add(ins)
	}

	// total统计
	set.Total = len(set.Items)

	return set, nil
}

func (i *HostServiceImpl) QueryHost(ctx context.Context, req *host.QueryHostRequest) (*host.Host, error) {
	args := []interface{}{req.Xid}
	i.log.Named("QueryHost").Infof("query sql: %s; %v", QueryHostSQL, args)

	// query stmt, 构建一个Prepare语句
	stmt, err := i.db.PrepareContext(ctx, QueryHostSQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// 取出数据，赋值结构体
	ins := host.NewHost()
	err = stmt.QueryRowContext(ctx, args...).Scan(&ins.Xid, &ins.Status, &ins.CreateAt, &ins.CreateBy, &ins.UpdateAt, &ins.UpdateBy, &ins.DeleteAt, &ins.DeleteBy, &ins.ResourceId, &ins.Vendor,
		&ins.Name, &ins.Region, &ins.ExpireAt, &ins.PublicIP, &ins.PrivateIP, &ins.CPU, &ins.Memory, &ins.OSType, &ins.OSName, &ins.SerialNumber,
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

	resStmt, err := tx.PrepareContext(ctx, updateHostSQL)
	if err != nil {
		return nil, err
	}
	_, err = resStmt.ExecContext(ctx, ins.UpdateAt, ins.UpdateBy, ins.Vendor, ins.Name, ins.PublicIP, ins.PrivateIP, ins.CPU, ins.Memory, ins.OSType, ins.OSName, ins.SerialNumber)
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
	_, err = resStmt.ExecContext(ctx, req.Xid)
	if err != nil {
		return err
	}

	return nil
}
