import React from "react";
import {Button, Table} from "antd";
import BlankElement from "../components/BlankElement";
import HttpClient from "../components/HttpClient";
import HttpURL from "../env/HttpURL";

class State {
    dataSource = []
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

    render() {
        return <Table columns={this.columns} dataSource={this.state.dataSource}/>;
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