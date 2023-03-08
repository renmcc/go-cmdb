package impl

const (
	InsertHostSQL = `INSERT INTO host(xid,status,create_at,create_by,resource_id,vendor,name,region,expire_at,public_ip,private_ip,cpu,memory,os_type,os_name,serial_number) 
					VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`

	ListHostSQL = `SELECT xid,status,create_at,create_by,update_at,update_by,delete_at,delete_by,resource_id,vendor,name,region,expire_at,public_ip,private_ip,cpu,memory,os_type,os_name,serial_number 
		FROM host 
	  		WHERE STATUS = 1 AND private_ip LIKE ? AND serial_number LIKE ? LIMIT ?,?`

	QueryHostSQL = `SELECT xid,status,create_at,create_by,update_at,update_by,delete_at,delete_by,resource_id,vendor,name,region,expire_at,public_ip,private_ip,cpu,memory,os_type,os_name,serial_number 
		FROM host 
	  		WHERE STATUS = 1 AND xid = ?`

	updateHostSQL = `UPDATE host SET update_at=?,update_by=?,vendor=?,name=?,public_ip=?,private_ip=?,cpu=?,memory=?,os_type=?,os_name=?,serial_number=?
	WHERE status=1 AND xid = ?`

	deleteHostSQL = "UPDATE host SET status=0 WHERE xid = ?"
)
