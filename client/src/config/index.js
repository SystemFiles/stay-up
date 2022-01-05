import * as dotenv from 'dotenv'
import { DEVELOPMENT, PRODUCTION, TEST } from '../common/constants'

// Import env vars from appropriate file
switch (process.env.NODE_ENV) {
	case DEVELOPMENT:
		dotenv.config({ path: '.dev.env' })
		break
	case TEST:
		dotenv.config({ path: '.test.env' })
	case PRODUCTION:
		dotenv.config({ path: '.prod.env' })
	default:
		dotenv.config()
		break
}

export default {
	apiBaseUrl : process.env.API_BASE_URL || 'http://localhost:5555/api'
}
