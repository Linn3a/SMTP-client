import React from 'react';
import axios from 'axios';
import icon from '../assets/icon.png';
import {
    LockOutlined,
    UserOutlined,
  } from '@ant-design/icons';
  import {
    LoginForm,
    ProFormCheckbox,
    ProFormText,
    ProConfigProvider,
  } from '@ant-design/pro-components';
  import { notification } from 'antd';
import { useNavigate } from 'react-router-dom';
  axios.defaults.baseURL = 'http://localhost:8080/';
const Login = ({setIslogin}) => {
    const navigate = useNavigate();
    return (
        <ProConfigProvider hashed={false}>
          <div style={{ backgroundColor: 'white' }}>
            <LoginForm
              logo={icon}
              title="邮件客户端"
              subTitle="请登录"
              onFinish={async (values) => {
                console.log(values);
                const {data} = await axios.post('login',{
                    username:values.username,
                    password:values.password,
                  })
                  console.log(data);
                  if(data.status === "ok") {
                    setIslogin(true);
                    notification.success({message:"登录成功"})
                  }
                  else {
                    notification.error({message:"登录失败"})
                  }
        
              }}
              
              >
  
              <ProFormText
                name="username"
                fieldProps={{
                  size: 'large',
                  prefix: <UserOutlined className={'prefixIcon'} />,
                }}
                placeholder={'用户名'}
                rules={[
                  {
                    required: true,
                    message: '请输入用户名!',
                  },
                ]}
              />
              <ProFormText.Password
                name="password"
                    fieldProps={{
                  size: 'large',
                  prefix: <LockOutlined className={'prefixIcon'} />,
                }}
                placeholder={'密码'}
                rules={[
                  {
                    required: true,
                    message: '请输入密码！',
                  },
                ]}
              />
              <div
                style={{
                  marginBlockEnd: 24,
                }}
              >
                <ProFormCheckbox noStyle name="autoLogin">
                  自动登录
                </ProFormCheckbox>
              </div>
            </LoginForm>
          </div>
        </ProConfigProvider>
      );
}


export default Login;

  