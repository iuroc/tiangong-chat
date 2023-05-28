import { RouteEvent } from 'apee-router'

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
        if (phone.match(/^\s+$/) || password.match(/^\s+$/))
            return alert('输入不能为空')
        const xhr = new XMLHttpRequest()
        xhr.open('POST', 'https://neice.tiangong.cn/api/v1/user/login')
        const data = { data: { phone, passwd: password } }
        xhr.send(JSON.stringify(data))
    })
}