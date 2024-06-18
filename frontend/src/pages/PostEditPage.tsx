import type { FC } from "react";

import { Button } from "~/shared/components/Button";
import { Container } from "~/shared/components/Container";

export const PostEditPage: FC = () => {
  return (
    <Container>
      <form>
        <label>
          タイトル
          <input type="text" name="title" />
        </label>
        <label>
          本文
          <input type="text" name="body" />
        </label>

        <Button type="submit">更新する</Button>
      </form>
    </Container>
  );
};
