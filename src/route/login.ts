import { RouteEvent } from 'apee-router'
import { apiConfig } from '../config'
import { HttpResponse } from '../util'

/** `hash = '#/login'` */
export const login: RouteEvent = (route, router) => {
    if (route.status == 1) return
    /** 按钮：点击登录 */
    const loginBtn = route.dom.querySelector('.login') as HTMLButtonElement
    /** 输入框：手机号 */
    const phoneInput = route.dom.querySelector('.input-phone') as HTMLInputElement
    /** 输入框：密码 */
    const passwordInput = route.dom.querySelector('.input-password') as HTMLInputElement
    loginBtn.addEventListener('click', () => {
        let phone = phoneInput.value
        let password = passwordInput.value
        if (phone.match(/^\s*$/) || password.match(/^\s*$/))
            return alert('输入不能为空')
        const xhr = new XMLHttpRequest()
        xhr.open('POST', apiConfig.login)
        const postParam = new URLSearchParams()
        postParam.set('phone', phone)
        postParam.set('password', password)
        xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded')
        xhr.send(postParam.toString())
        xhr.addEventListener('readystatechange', () => {
            if (xhr.status == 200 && xhr.readyState == xhr.DONE) {
                const res: HttpResponse = JSON.parse(xhr.responseText)
                if (res.code == 200) {
                    let token = res.data as string
                    localStorage.setItem('token', token)
                    location.hash = ''
                    return
                }
                alert(res.msg)
            }
        })
    })
}

/** 登录校验 */
export function checkLoginInfo() {
    // 判断 token 和 invite-token 是否存在
    let token = localStorage.getItem('token')
    let inviteToken = localStorage.getItem('invite-token')
    if (!token) return location.hash = '/login'
}