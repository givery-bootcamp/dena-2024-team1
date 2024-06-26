import { Link } from "react-router-dom";

import { formatDateTime } from "~/shared/utils";
import { Post } from "~/generated";

type PostItemProps = {post: Post};

export function PostItem({ post }: PostItemProps) {
  return(
    <Link to={`/posts/${post.id}`} className="flex items-center justify-between">
      <div className="text-xl font-bold">
        {post.title}
      </div>
      <div>
        <div className='text-sm font-semibold text-gray-100'>
          {post.user_name}
        </div>
        <div className='text-sm text-gray-200'>
          更新日時: {formatDateTime(post.updated_at)}
        </div>
      </div>
    </Link>
  );
}
