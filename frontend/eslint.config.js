import globals from "globals";
import pluginJs from "@eslint/js";
import tseslint from "typescript-eslint";
import pluginReactConfig from "eslint-plugin-react/configs/recommended.js";
import { fixupConfigRules } from "@eslint/compat";
import tailwind from "eslint-plugin-tailwindcss"


export default [
  {languageOptions: { globals: globals.browser }},
  pluginJs.configs.recommended,
  ...tseslint.configs.recommended,
  ...fixupConfigRules(pluginReactConfig),
  ...tailwind.configs["flat/recommended"],
  {
    rules: {
      "react/jsx-uses-react": "off",
      "react/react-in-jsx-scope": "off"
    },
    settings: {
      react: {
        version: 'detect',
      },
    },
  }
];
