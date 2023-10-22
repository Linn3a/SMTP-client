import React from 'react';
import { Tag, Button } from 'antd';
import { PageContainer,ProTable } from '@ant-design/pro-components';
import { fetchAllCrafts } from './api';
import { useQuery } from '@tanstack/react-query';
import { useNavigate } from 'react-router-dom';



const Box  = () => {
    const navigate = useNavigate();
    const {data: crafts} = useQuery({ queryKey: ['crafts'], queryFn: fetchAllCrafts });
      console.log(crafts);
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
      {
        title: '操作',
        valueType: 'option',
        render: (_1, record, _, action) => [
          <Button onClick={() =>{
            navigate(`${record.ID}`)
          }}>编辑草稿</Button>
        ]
      },
        
];
  return (
    <PageContainer
      header={{
        title:"草稿箱"
      }}
      >
      
        <ProTable
            search={false}
            rowKey="ID"
            dataSource={crafts}
            columns={columns}
        />
    </PageContainer>
  );
}


export default Box;