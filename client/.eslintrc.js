// https://eslint.org/docs/user-guide/configuring

module.exports = {
  root: true,
  parserOptions: {
    parser: 'babel-eslint'
  },
  env: {
    browser: true,
  },
  extends: [
    // https://github.com/vuejs/eslint-plugin-vue#priority-a-essential-error-prevention
    // consider switching to `plugin:vue/strongly-recommended` or `plugin:vue/recommended` for stricter rules.
    'plugin:vue/essential',
    // https://github.com/standard/standard/blob/master/docs/RULES-en.md
    'standard'
  ],
  // required to lint *.vue files
  plugins: [
    'vue'
  ],
  // add your custom rules here
  rules: {
	  'no-new': 0,
	  'no-tabs':0,
	  'no-mixed-spaces-and-tabs':0,
	  'semi':0,
	  'one-var':0,
	  'no-extra-boolean-cast':0,
	  'object-curly-spacing':0,
	  'promise/param-names': 0,
	  // allow paren-less arrow functions
	  'arrow-parens': 0,
	  // allow debugger during development
	  'no-debugger': process.env.NODE_ENV === 'production' ? 2 : 0,
	  'indent': 0,
	  'space-before-function-paren': 0,
	  'eol-last': 0,
	  'vue/max-attributes-per-line':0,
	  'vue/attribute-hyphenation': 'off',//关闭连字符
	  'vue/html-self-closing':0,
	  'vue/html-indent':0,
	  'no-useless-escape': 0,
	  'no-undef': 0,
	  "no-unused-expressions": 0,//禁止无用的表达式
}
}
