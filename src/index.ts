import { Router } from 'apee-router'
import { checkLoginInfo, login } from './route/login'

export const router = new Router()
router.set(['home', 'login'])
router.set('login', login)
router.start()
checkLoginInfo()