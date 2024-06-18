import { Button } from "~/shared/components/Button";
import { Container } from "~/shared/components/Container";

export function PostDetailPage() {
  return (
    <Container>
      <div className="flex flex-col gap-6">
        <div className="mt-10 flex flex-col gap-7">
          <div className="flex flex-col gap-2">
            <h1 className="text-xl font-bold">{mockPost.title}</h1>
            <div className="flex flex-col gap-0.5 text-sm text-gray-200">
              <p>作成日時: {mockPost.createdAt}</p>
              <p>更新日時: {mockPost.updatedAt}</p>
            </div>
          </div>
          <hr className="border-border" />
          <div className="flex flex-col gap-3 text-lg">
            <p className="whitespace-pre-line leading-8">{mockPost.body}</p>
            <p className="text-right">{mockPost.user.username}</p>
          </div>
        </div>
        <div className="flex gap-2">
          <Button>編集</Button>
          <Button variant="alert">削除</Button>
        </div>
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
