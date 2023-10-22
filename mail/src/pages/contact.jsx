import React from 'react';
import { PageContainer,ProTable } from '@ant-design/pro-components';
import { fetchAllUsers } from './api';
import { useQuery } from '@tanstack/react-query';



const Contact  = () => {
    const {data: users } = useQuery({ queryKey: ['users'], queryFn: fetchAllUsers });
      console.log(users);
    const columns = [{
        title: 'ID',
        dataIndex: 'ID',
        },{
        title: '用户名',
        dataIndex: 'Username',
        },
        {
        title: '地址',
        dataIndex:'Address',
    },
        
];
  return (
    <PageContainer
      header={{
        title:"通讯录"
      }}
      >
      
        <ProTable
            search={false}
            rowKey="ID"
            dataSource={users}
            columns={columns}
        />
    </PageContainer>
  );
}


export default Contact;