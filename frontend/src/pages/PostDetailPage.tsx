import { useEffect } from "react";
import { useParams } from "react-router-dom";

import { useAppDispatch, useAppSelector } from "~/shared/hooks";
import { APIService } from "~/shared/services";
import { Container } from "~/shared/components/Container";
import { formatDateTime } from "~/shared/utils";

export function PostDetailPage() {
  const { post } = useAppSelector((state) => state.post);
  const dispatch = useAppDispatch();
  const { postId } = useParams<{postId: string}>();
  useEffect(() => {
    dispatch(APIService.getPost(Number(postId)));
  }, [dispatch]);
  return (
    <Container>
      <div className="mt-10 flex flex-col gap-7">
        <div className="flex flex-col gap-2">
          <h1 className="text-xl font-bold">{post?.title}</h1>
          <div className="flex flex-col gap-0.5 text-sm text-gray-200">
            <p>作成日時: { formatDateTime(post?.createdAt ?? "") }</p>
            <p>更新日時: { formatDateTime(post?.updatedAt ?? "")}</p>
          </div>
        </div>
        <hr className="border-border" />
        <div className="flex flex-col gap-3 text-lg">
          <p className="whitespace-pre-line leading-8">{post?.body}</p>
          <p className="text-right">{post?.userName}</p>
        </div>
      </div>
    </Container>
  );
}
