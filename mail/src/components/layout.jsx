import React from 'react';
import {  ProLayout } from '@ant-design/pro-components';
import { Outlet,NavLink } from 'react-router-dom';
import { TeamOutlined,EditOutlined,MailOutlined,SnippetsOutlined } from '@ant-design/icons';


const route = {
    path: '/',
    routes: [
      {
        path: '/contact/',
        name: '通讯录',
        icon: <TeamOutlined />,
      },
      {
        path: '/send/',
        name: '发邮箱',
        icon: <EditOutlined />,
      },
      {
        path: '/box/',
        name: '已发送',
        icon: <MailOutlined />,
      },
      {
        path:'/craft/',
        name:'草稿箱',
        icon:<SnippetsOutlined  />,
      },
    ],
  }

  
const Layout  = () => {
  return (
    <ProLayout
        title = "邮件客户端"
        locale='zh-CN'
        logo="https://em-content.zobj.net/source/microsoft/319/e-mail_1f4e7.png"
        route={route}
        menuItemRender={(item, dom) => 
          <NavLink to={item.path || '/'}>{dom}</NavLink>
        }
    >
      <Outlet/>
    </ProLayout>
  );
}


export default Layout;