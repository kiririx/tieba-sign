export default class StringUtil {
    static hasText(value: string): boolean {
        return value !== undefined && value !== null && value.trim() != "";
    }

    static isBlank(value: string): boolean {
        return !this.hasText(value);
    }

    static contains(value: string, containVal: string): boolean {
        return value.indexOf(containVal) > -1;
    }
}