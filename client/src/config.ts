export default {
  isProduction: process.env.APP_ENV === 'production',
  hostname: process.env.HOSTNAME || 'https://127.0.0.1:50051',
}
