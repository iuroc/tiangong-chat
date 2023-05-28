import { Router } from 'apee-router'
import { checkLoginInfo } from './util'
import { login } from './route/login'

const router = new Router()
router.set(['home', 'login'])
router.set('login', login)
router.start()
checkLoginInfo()