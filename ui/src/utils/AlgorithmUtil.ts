import Base64 from './base64';
let md5 = require('js-md5');

export default class AlgorithmUtil {

    public static encodeBase64(text: string): string {
        return Base64.encode(text);
    }

    public static decodeBase64(text: string): string {
        return Base64.decode(text);
    }

    public static URLEncoding(text: string) {
        return encodeURIComponent(text);
    }

    public static URLDecoding(text: string) {
        return decodeURIComponent(text);
    }

    public static encodeMD5(text: string) {
        return md5(text);
    }

    public static unicodeToZh(text: string) {
        let _text = text.replace(/\\/g, "%");
        return unescape(_text);
    }

    public static zhToUnicode(text: string) {
        let res = [];
        for (let i = 0; i < text.length; i++) {
            res[i] = ( "00" + text.charCodeAt(i).toString(16) ).slice(-4);
        }
        return "\\u" + res.join("\\u");
    }

}