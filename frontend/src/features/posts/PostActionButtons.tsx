import type { FC } from "react";
import { Link } from "react-router-dom";

import { Button } from "~/shared/components/Button";

type Props = {
  postId: number;
  onDelete: () => void;
}

export const PostActionButtons: FC<Props> = ({ postId, onDelete }) => {
  return (
    <div className="flex gap-2">
      <Link to={`/posts/${postId}/edit`}>
        <Button>編集</Button>
      </Link>
      <Button variant="alert" onClick={onDelete}>削除</Button>
    </div>
  );
};
