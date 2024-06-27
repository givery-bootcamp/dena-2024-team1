export type PageType = "sign-in" | "sign-up";

export const typeToTitle = {
  "sign-in": "ログイン",
  "sign-up": "新規登録",
} as const;
export const typeToButtonLabel = {
  "sign-in": "ログインする",
  "sign-up": "アカウントを作成する",
} as const;
export const typeToLink = {
  "sign-in": "/signup",
  "sign-up": "/signin",
} as const;
export const typeToNavigationText = {
  "sign-in": "アカウントをお持ちでない方はこちら",
  "sign-up": "アカウントをお持ちの方はこちら",
} as const;
