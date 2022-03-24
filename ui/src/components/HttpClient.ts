import axios from "axios";
import ObjectUtil from "../utils/ObjectUtil";
import {notification} from "antd";
import Resp from "../model/HttpModel";
import RouterURL from "../env/RouterURL";
import RouterUtil from "../utils/RouterUtil";

axios.interceptors.response.use(
    (response: any) => {
        const pathName = RouterUtil.getPath();
        if (!response.data.login && pathName != RouterURL.LOGIN) {
            RouterUtil.push(RouterURL.LOGIN);
        } else {
            return response;
        }
    }
);

export default class HttpClient {
    public static post = (url: string, data: {} = {}, successFunc?: (resp: Resp) => void, finalFunc?: (err?: any) => void): void => {
        axios.post(url, data).then(value => {
            if (ObjectUtil.isNotNull(value)) {
                let _data = value.data;
                if (_data.status === 'error') {
                    console.log('errLog[' + url + ']:' + _data.data.errormsg);
                    notification.open({
                        message: '错误信息',
                        description: _data.data.errormsg,
                        duration: 1.5
                    });
                } else {
                    if (successFunc) {
                        successFunc(_data as Resp);
                    }
                }
            }
            if (finalFunc) {
                finalFunc();
            }
        }).catch(err => {
            console.log(err);
            if (finalFunc) {
                finalFunc(err);
            }
        });
    }

    public static get = (url: string, successFunc?: (resp: Resp) => void, finalFunc?: (err?: any) => void): void => {
        axios.get(url).then(value => {
            if (ObjectUtil.isNotNull(value)) {
                let _data = value.data;
                if (_data.status === 'error') {
                    console.log('errLog[' + url + ']:' + _data.data.errormsg);
                    notification.open({
                        message: '错误信息',
                        description: _data.data.errormsg,
                        duration: 1.5
                    });
                } else {
                    if (successFunc) {
                        successFunc(_data as Resp);
                    }
                }
            }
            if (finalFunc) {
                finalFunc();
            }
        }).catch(err => {
            console.log(err);
            if (finalFunc) {
                finalFunc(err);
            }
        });
    }

    public static put = (url: string, data: {} = {},successFunc?: (resp: Resp) => void, finalFunc?: (err?: any) => void): void => {
        axios.put(url, data).then(value => {
            if (ObjectUtil.isNotNull(value)) {
                let _data = value.data;
                if (_data.status === 'error') {
                    console.log('errLog[' + url + ']:' + _data.data.errormsg);
                    notification.open({
                        message: '错误信息',
                        description: _data.data.errormsg,
                        duration: 1.5
                    });
                } else {
                    if (successFunc) {
                        successFunc(_data as Resp);
                    }
                }
            }
            if (finalFunc) {
                finalFunc();
            }
        }).catch(err => {
            console.log(err);
            if (finalFunc) {
                finalFunc(err);
            }
        });
    }

    public static delete = (url: string, successFunc?: (resp: Resp) => void, finalFunc?: (err?: any) => void): void => {
        axios.delete(url).then(value => {
            if (ObjectUtil.isNotNull(value)) {
                let _data = value.data;
                if (_data.status === 'error') {
                    console.log('errLog[' + url + ']:' + _data.data.errormsg);
                    notification.open({
                        message: '错误信息',
                        description: _data.data.errormsg,
                        duration: 1.5
                    });
                } else {
                    if (successFunc) {
                        successFunc(_data as Resp);
                    }
                }
            }
            if (finalFunc) {
                finalFunc();
            }
        }).catch(err => {
            console.log(err);
            if (finalFunc) {
                finalFunc(err);
            }
        });
    }
}