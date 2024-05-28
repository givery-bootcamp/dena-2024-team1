import globals from "globals";
import pluginJs from "@eslint/js";
import tseslint from "typescript-eslint";
import pluginReactConfig from "eslint-plugin-react/configs/recommended.js";
import { fixupConfigRules } from "@eslint/compat";
import tailwind from "eslint-plugin-tailwindcss";
import importPlugin from "eslint-plugin-import";
import noRelativeImportPaths from "eslint-plugin-no-relative-import-paths";

export default [
  { languageOptions: { globals: globals.browser } },
  pluginJs.configs.recommended,
  ...tseslint.configs.recommended,
  ...fixupConfigRules(pluginReactConfig),
  ...tailwind.configs["flat/recommended"],
  {
    plugins: {
      import: importPlugin,
      "no-relative-import-paths": noRelativeImportPaths,
    },
  },
  {
    rules: {
      "react/jsx-uses-react": "off",
      "react/react-in-jsx-scope": "off",
      "react/jsx-indent" : ["error", 2],
      "semi": "error",
      "quotes": "error",
      "comma-dangle": ["error", "always-multiline"],
      "object-curly-spacing": ["error", "always"],
      "eol-last": "error",
      "import/order": [
        "error",
        {
          groups: ["builtin", "external", "internal", "sibling", "parent", "index", "object"],
          "newlines-between": "always",
        },
      ],
      "no-relative-import-paths/no-relative-import-paths": [
        "error",
        { "allowSameFolder": true, "rootDir": "src", "prefix": "~" },
      ],
    },
    settings: { react: { version: "detect" } },
  },
];
