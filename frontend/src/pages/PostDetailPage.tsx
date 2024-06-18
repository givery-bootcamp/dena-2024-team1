import { PostActionButtons } from "~/features/posts/PostActionButtons";
import { PostContent } from "~/features/posts/PostContent";
import { PostHeading } from "~/features/posts/PostHeading";
import { Container } from "~/shared/components/Container";

export function PostDetailPage() {
  return (
    <Container>
      <div className="flex flex-col gap-6">
        <div className="mt-10 flex flex-col gap-7">
          <PostHeading
            title={mockPost.title}
            createdAt={mockPost.createdAt}
            updatedAt={mockPost.updatedAt}
          />
          <hr className="border-border" />
          <PostContent body={mockPost.body} username={mockPost.user.username} />
        </div>
        <PostActionButtons postId={mockPost.id} />
      </div>
    </Container>
  );
}

// APIから以下のレスポンスが返ってくる想定
const mockPost = {
  id: "1",
  title: "バナナはおやつに含まれますか？",
  body: "来週の遠足の件で質問です。\n\nおやつは一人500円までと言われているのですが、バナナはおやつに含まれますか？\nそれとも弁当に含まれますか？",
  createdAt: "2024/05/27 09:00",
  updatedAt: "2024/05/27 15:30",
  user: {
    username: "watapon",
  },
};
