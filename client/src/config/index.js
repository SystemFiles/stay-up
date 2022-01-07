import * as dotenv from 'dotenv'
dotenv.config()

// Set variables
export const API_BASE_URL = process.env.API_BASE_URL || 'http://localhost:5555/api'
export const API_WEBSOCK_URL = process.env.API_WEBSOCK_URL || 'ws://localhost:5555/api/service/ws'
