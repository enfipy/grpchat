module.exports = {
  presets: [
    [
      '@babel/env',
      {
        modules: 'commonjs',
        targets: {
          browsers: ['> 1%', 'last 2 versions', 'not ie <= 8']
        }
      }
    ],
    '@vue/app'
  ],
  plugins: ['add-module-exports']
}
