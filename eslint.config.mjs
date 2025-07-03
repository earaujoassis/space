import { defineConfig } from "eslint/config";
import react from "eslint-plugin-react";
import prettier from "eslint-config-prettier";
import globals from "globals";
import path from "node:path";
import { fileURLToPath } from "node:url";
import js from "@eslint/js";
import { FlatCompat } from "@eslint/eslintrc";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);
const compat = new FlatCompat({
    baseDirectory: __dirname,
    recommendedConfig: js.configs.recommended,
    allConfig: js.configs.all
});

export default defineConfig([{
    files: ['**/*.js', '**/*.jsx'],

    extends: compat.extends("eslint:recommended", "plugin:react/recommended", "prettier"),

    plugins: {
        react,
    },

    languageOptions: {
        globals: {
            ...globals.node,
            ...globals.browser,
            NProgress: "readonly",
            __COMMIT_HASH__: "readonly",
        },

        ecmaVersion: 2018,
        sourceType: "module",

        parserOptions: {
            ecmaFeatures: {
                jsx: true,
                experimentalObjectRestSpread: true,
            },
        },
    },

    settings: {
        react: {
            version: "detect",
        },
    },

    rules: {
        indent: ["error", 2, {
            SwitchCase: 1,
        }],

        "react/jsx-indent": ["error", 2],
        "react/jsx-indent-props": ["error", 2],
        "@/jsx-quotes": ["error", "prefer-double"],

        "linebreak-style": ["error", "unix"],
        quotes: ["error", "single"],
        semi: ["error", "never"],
        "array-bracket-spacing": ["error", "never"],
        "react/prop-types": 0,

        ...prettier.rules
    },
}]);
