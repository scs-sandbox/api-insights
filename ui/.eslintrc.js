module.exports = {
  env: {
    browser: true,
    es2021: true,
  },
  extends: [
    'airbnb',
    'airbnb/hooks',
    // Disable 'react/react-in-jsx-scope' and 'react/jsx-uses-react' rules for React over 17
    'plugin:react/jsx-runtime',
    'plugin:@typescript-eslint/recommended',
  ],
  overrides: [
  ],
  parser: '@typescript-eslint/parser',
  plugins: [
    '@typescript-eslint',
  ],
  settings: {
    // Use tsconfig.json as module resolver
    'import/resolver': {
      typescript: true,
      javascript: true,
    },
  },
  rules: {
    // Customized for typescript file extensions
    'import/extensions': [
      'error',
      'ignorePackages',
      {
        js: 'never',
        jsx: 'never',
        ts: 'never',
        tsx: 'never',
      },
    ],
    // Customized for typescript components
    'react/jsx-filename-extension': ['error', { extensions: ['.jsx', '.tsx'] }],
    // Customized for enum by 2 following items
    'no-shadow': 'off',
    '@typescript-eslint/no-shadow': ['error'],
    //
    '@typescript-eslint/no-namespace': 'off',
    //
    'jsx-a11y/click-events-have-key-events': 'off',
    'jsx-a11y/no-static-element-interactions': 'off',
    'jsx-a11y/label-has-associated-control': 'off',
    'jsx-a11y/no-autofocus': 'off',
    'react/destructuring-assignment': 'off',
    'react-hooks/exhaustive-deps': 'off',
    'react/jsx-props-no-spreading': 'off',
    'react/require-default-props': 'off',
  },
};
