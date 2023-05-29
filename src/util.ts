/** HTTP 响应 */
export type HttpResponse = {
    /** 响应码 */
    code: number,
    /** 响应提示 */
    msg: string,
    /** 响应数据 */
    data: any
}