import { Container } from "~/shared/components/Container";

export function CreatePostPage() {
  return (
    <Container className="mt-10">
      <div className="rounded-lg bg-white p-8 shadow-md">
        <form>
          <div>
            <label>
              タイトル
            </label>
            <br/>
            <input type="text" value="" />
          </div>
          <br/>
          <div>
            <label>
              本文
            </label>
            <br/>
            <textarea value="" />
          </div>
        </form>
      </div>
    </Container>
  );
}
