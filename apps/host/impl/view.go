package impl

import (
	"context"

	"github.com/renmcc/go-cmdb/apps/host"
)

func (i *HostServiceImpl) CreateHost(ctx context.Context, ins *host.Host) (*host.Host, error) {
	// 校验数据合法性
	if err := ins.Validate(); err != nil {
		return nil, err
	}

	// 默认值填充
	ins.InjectDefault()

	// 对象入库
	if err := i.save(ctx, ins); err != nil {
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
		if err := rows.Scan(
			&ins.Id, &ins.Vendor, &ins.Region, &ins.CreateAt, &ins.ExpireAt,
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

func (i *HostServiceImpl) DescribeHost(ctx context.Context, req *host.QueryHostRequest) (*host.Host, error) {
	return nil, nil
}

func (i *HostServiceImpl) UpdateHost(ctx context.Context, req *host.UpdateHostRequest) (*host.Host, error) {
	return nil, nil
}

func (i *HostServiceImpl) DeleteHost(ctx context.Context, req *host.DeleteHostRequest) (*host.Host, error) {
	return nil, nil
}
