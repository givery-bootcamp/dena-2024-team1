export type PageType = "sign-in" | "sign-up";

type Label = "ログイン" | "新規登録";
type PageLink = "/signin" | "/signup";
type ButtonLabel = "ログインする" | "アカウントを作成する";

export const typeToLabel: Record<PageType, Label> = {
  "sign-in": "ログイン",
  "sign-up": "新規登録",
};
export const typeToButtonLabel: Record<PageType, ButtonLabel> = {
  "sign-in": "ログインする",
  "sign-up": "アカウントを作成する",
};
export const typeToLink: Record<PageType, PageLink> = {
  "sign-in": "/signup",
  "sign-up": "/signin",
};
export const typeToNavigationText: Record<PageType, string> = {
  "sign-in": "アカウントをお持ちでない方はこちら",
  "sign-up": "アカウントをお持ちの方はこちら",
};
