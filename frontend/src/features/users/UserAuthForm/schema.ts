import * as v from "valibot";

export const schema = v.object({
  username: v.pipe(
    v.string(),
    v.nonEmpty("ユーザー名を入力してください"),
  ),
  password: v.pipe(
    v.string(),
    v.nonEmpty("パスワードを入力してください"),
    v.minLength(8, "パスワードは8文字以上で入力してください"),
  ),
});

export type Schema = v.InferOutput<typeof schema>
