import { Container } from "../shared/components/Container"

export function PostDetailPage() {
  return (
    <Container>
      <h1>{mockPost.title}</h1>
      <div>
        <p>作成日時: {mockPost.createdAt}</p>
        <p>更新日時: {mockPost.updatedAt}</p>
      </div>
      <hr />
      <div>
        <p>{mockPost.body}</p>
        <p>{mockPost.user.username}</p>
      </div>
    </Container>
  )
}

const mockPost = {
  id: "1",
  title: "バナナはおやつに含まれますか？",
  body: "来週の遠足の件で質問です。\nおやつは一人500円までと言われているのですが、バナナはおやつに含まれますか？\nそれとも弁当に含まれますか？",
  createdAt: "2024/05/27 09:00",
  updatedAt: "2024/05/27 15:30",
  user: {
    username: "watapon"
  }
}
