import React from 'react';
import { Tag } from 'antd';
import { PageContainer,ProTable } from '@ant-design/pro-components';
import { fetchAllMails } from './api';
import { useQuery } from '@tanstack/react-query';



const Box  = () => {
    const {data: mails } = useQuery({ queryKey: ['mails'], queryFn: fetchAllMails });
      console.log(mails);
    const columns = [{
        title: 'ID',
        dataIndex: 'ID',
        },
        {
            title: '发送时间',
            dataIndex: 'CreatedAt',
        },{
        title: '标题',
        dataIndex: 'Subject',
        },
        {
        title: '内容',
        dataIndex:'Content',
    },
    {
        title: '发送到',
        key: 'to',
        dataIndex: 'To',
        render: (_, { To }) => (
          <>
            {To.map((t) => {
              let color = 'geekblue';
            
              return (
                <Tag color={color} key={t}>
                  {t.Address}
                </Tag>
              );
            })}
          </>
        ),
      },
        
];
  return (
    <PageContainer
      header={{
        title:"已发送"
      }}
      >
      
        <ProTable
            search={false}
            rowKey="ID"
            dataSource={mails}
            columns={columns}
        />
    </PageContainer>
  );
}


export default Box;