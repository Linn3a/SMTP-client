import {
  ProForm,
  ProFormList,
  ProFormSelect,
  ProFormText,
  ProFormTextArea,
} from '@ant-design/pro-components';
import { message,Button } from 'antd';
import axios from 'axios';
import { useRef, useState } from 'react';

axios.defaults.baseURL = 'http://127.0.0.1:8080'



const Send = () => {
const [Users, setUsers] = useState()
  const formRef = useRef()
  const onSave = () => {
    const values = formRef.current.getFieldValue();

    // console.log(Users);
    let to = []
    values.users.map((u,index) => {
      console.log(u);
      console.log("Users", Users);
        to.push({
            username:Users[u].Username,
            address:Users[u].Address

        })
    })
    if (values.datas) {
     to = [...to, ...values.datas]
    }
    const body = {
        "subject":values.subject,
        "to":to,
        "content":values.content,
    }

    console.log("body", body);

    axios.post("/crafts", body).then(res => {
      console.log(res);
      if (res.status === 200) {
        message.success('保存成功');
      }
    })
  };
  return (
    <ProForm
      color='blue'
      
      onFinish={async (values) => {
        console.log(values);
        console.log(Users);
        let to = []
        values.users.map((u,index) => {
          console.log(u);
          console.log("Users", Users);
            to.push({
                username:Users[u].Username,
                address:Users[u].Address

            })
        })
        if (values.datas) {
         to = [...to, ...values.datas]
        }
        const body = {
            "subject":values.subject,
            "to":to,
            "content":values.content,
        }

        console.log("body", body);

        const data = await axios.post("/send", body)
        console.log("data", data)
        message.success('提交成功');
      }}
      submitter={{
        render: (props, doms) => {
          return [
            ...doms,
            <Button htmlType="button" onClick={onSave} key="craft">
              保存到草稿
            </Button>,
            
           
          ];
        },
      }}

      formRef={formRef}
      params={{ id: '100' }}
      formKey="base-form-use-demo"
    
      autoFocusFirstInput
    >
      <ProFormSelect
            label="发送到"
            name="users"
            mode="multiple"
            rules={[
              {
                required: true,
              },
            ]}
            request={() => axios.get('/users')
            .then(res => {
                console.log(res.data);
                setUsers(res.data.users)

               return  res.data.users.map((u) => {
                return {label: u.Address, value: u.id}})})}
          />

<ProFormList name="datas">
        {() => {
          return (
            <>
              <ProFormText
                label="新增用户用户名"
                name="username"
                placeholder="输入用户名"
              />
                <ProFormText
                label="新增用户邮箱"
                name="address"
                placeholder="输入邮箱地址"
              />
              </>
          );
        }}
      </ProFormList>
        
     
      <ProForm.Group>
        <ProFormText
          name="subject"
          width="md"
          label="邮件标题"
          placeholder="请输入标题"
        />
        
      </ProForm.Group>
        
      <ProFormTextArea
        colProps={{ span: 24 }}
        name="content"
        label="邮件内容"
      />
     
    </ProForm>
  );
};

export default Send