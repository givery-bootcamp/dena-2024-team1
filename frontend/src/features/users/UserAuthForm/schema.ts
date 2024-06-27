import * as v from "valibot";

export const schema = v.object({
  username: v.pipe(
    v.string(),
    v.nonEmpty("ユーザー名を入力してください"),
  ),
  password: v.pipe(
    v.string(),
    v.nonEmpty("パスワードを入力してください"),
  ),
});

export type Schema = v.InferOutput<typeof schema>
