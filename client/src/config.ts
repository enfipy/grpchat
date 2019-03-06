export default {
  isProduction: process.env.APP_ENV === 'production',
  hostname: process.env.HOSTNAME || 'http://localhost:50051',
}
