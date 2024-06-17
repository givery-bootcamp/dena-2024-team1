import { Link } from "react-router-dom";

import { Post } from "~/shared/models";

type PostItemProps = {post: Post};

export function PostItem({ post }: PostItemProps) {
  return(
    <Link to={`/posts/${post.id}`} className="flex items-center justify-between">
      <div className="text-xl font-bold">
        {post.title}
      </div>
      <div>
        <div className='text-sm font-semibold text-gray-100'>
          {post.userName}
        </div>
        <div className='text-sm text-gray-200'>
          更新日時: {formatDateTime(post.updateAt)}
        </div>
      </div>
    </Link>
  );
}

// バックエンドから受け取った時間をフォーマットする関数
export function formatDateTime(dateTime: string): string {
  const date = new Date(dateTime);
  return date.toLocaleString([], { year: "numeric", month: "2-digit", day: "2-digit", hour: "2-digit", minute: "2-digit" });
}
