import StringUtil from "./StringUtil";

export default class ObjectUtil {
    static isNull(val: any): boolean {
        return val == null || val == undefined;
    }

    static isNotNull(val: any): boolean {
        return !this.isNull(val);
    }

    static getNewProperty = (obj: { [key: string]: any }, key: string, val: any): any => {
        if (StringUtil.contains(key, ".")) {
            const firstKey: string = key.substring(0, key.indexOf("."));
            const nextKey: string = key.substring(key.indexOf(".") + 1);
            const nextObj: any = obj[firstKey];
            return ObjectUtil.getNewProperty(nextObj, nextKey, val);
        } else {
            let property: any = obj[key];
            if (ObjectUtil.isNotNull(property)) {
                obj[key] = val;
            }
        }
        return obj;
    }

    static getClearedObject = (obj: { [key: string]: any }) => {
        for (const key in obj) {
            const val = obj[key];
            if (val) {
                if (typeof val == "string") {
                    obj[key] = "";
                } else if (typeof val == "boolean") {
                    obj[key] = false;
                } else if (typeof val == "object") {
                    obj[key] = ObjectUtil.getClearedObject(val);
                } else if (typeof val == "number") {
                    obj[key] = 0;
                } else {
                    obj[key] = "";
                }
            }
        }
        return obj;
    }

}