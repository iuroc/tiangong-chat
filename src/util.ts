/** 登录校验 */
export function checkLoginInfo() {
    // 判断 token 和 invite-token 是否存在
    let token = localStorage.getItem('token')
    let inviteToken = localStorage.getItem('invite-token')
    if (!token) return location.hash = '/login'
}

/** HTTP 响应 */
export type HttpResponse = {
    code: number,
    msg: string,
    data: any
}