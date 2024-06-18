import type { FC } from "react";
import { Link } from "react-router-dom";

import { Button } from "~/shared/components/Button";

type Props = {
  postId: string;
}

export const PostActionButtons: FC<Props> = ({ postId }) => {
  return (
    <div className="flex gap-2">
      <Link to={`/posts/${postId}/edit`}>
        <Button>編集</Button>
      </Link>
      <Button variant="alert">削除</Button>
    </div>
  );
};
