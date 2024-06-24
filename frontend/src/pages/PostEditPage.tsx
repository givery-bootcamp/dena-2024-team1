import type { FC } from "react";

import { PostEditForm } from "~/features/posts/PostEditForm";
import { Container } from "~/shared/components/Container";

export const PostEditPage: FC = () => {
  return (
    <Container className="pt-10">
      <PostEditForm />
    </Container>
  );
};
