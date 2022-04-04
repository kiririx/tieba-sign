import React from "react";
import {Button, Form, Input, Layout, Modal, Table} from "antd";
import BlankElement from "../components/BlankElement";
import HttpClient from "../components/HttpClient";
import HttpURL from "../env/HttpURL";
import {Content, Header} from "antd/es/layout/layout";

class State {
    dataSource = []
    configVisible = false
}

class BdussList extends React.Component<any, State> {

    constructor(props: any) {
        super(props);
        this.state = new State()
    }


    private columns = [
        {
            title: '唯一标识',
            dataIndex: 'name',
            key: 'name'
        },
        {
            title: 'BDUSS',
            dataIndex: 'bduss',
            key: 'bduss'
        },
        {
            title: '签到状态',
            dataIndex: 'signStatus',
            key: 'signStatus'
        },
        {
            title: '签到次数',
            dataIndex: 'signCount',
            key: 'signCount'
        },
        {
            title: '操作',
            dataIndex: 'ope',
            key: 'ope',
            render: () => {
                return <div>
                    <Button type={"primary"}>手动签到</Button>
                    <BlankElement width={5}/>
                    <Button>编辑配置</Button>
                    <BlankElement width={5}/>
                    <Button type={"primary"} danger>删除</Button>
                </div>
            }
        }
    ]

    handleOK = () => {

        this.handleCancel()
    }

    handleCancel = () => {
        this.setState({
            configVisible: false
        })
    }

    handleOpenConfig = () => {
        this.setState({
            configVisible: true
        })
    }

    render() {
        return <div>
            <Modal title={"配置信息"}
                   visible={this.state.configVisible}
                   onOk={this.handleOK}
                   onCancel={this.handleCancel}
            >
                <Form
                    name="basic"
                    labelCol={{ span: 5 }}
                    wrapperCol={{ span: 16 }}
                    autoComplete="off"
                >
                    <Form.Item
                        label="用户标识"
                        name="username"
                        rules={[{ required: true, message: '请输入用户标识' }]}
                    >
                        <Input placeholder={"用户唯一标识"}/>
                    </Form.Item>
                    <Form.Item
                        label={"BDUSS"}
                        name="bduss"
                        rules={[{ required: true, message: '请输入BDUSS' }]}>
                        <Input.TextArea rows={4} placeholder={"登入电脑端贴吧，打开开发者工具F12, 查看Cookie里面的BDUSS值，并粘贴到此"}/>
                    </Form.Item>
                </Form>
            </Modal>
            <Layout>
                <Header>
                    <div>
                        <Button type={"primary"} onClick={this.handleOpenConfig}>添加配置</Button>
                    </div>
                </Header>
                <Content>
                    <Table columns={this.columns} dataSource={this.state.dataSource}/>
                </Content>
            </Layout>
        </div>;
    }

    componentDidMount() {
        HttpClient.get(HttpURL.BDUSS, resp => {
            this.setState({
                dataSource: resp.data.list.map((v: any) => {
                    if (v.bduss.length > 30) {
                        v.bduss = v.bduss.substring(0, 30) + "......"
                    }
                    return v
                })
            })
        })
    }
}

export default BdussList