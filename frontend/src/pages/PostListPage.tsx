import { useEffect,Fragment } from "react";

import { useAppDispatch, useAppSelector } from "~/shared/hooks";
import { APIService } from "~/shared/services";
import { Container } from "~/shared/components/Container";
import { PostItem } from "~/features/posts/PostItem";
import { BorderLine } from "~/shared/components/BorderLine";

export function PostListPage() {
  const { posts } = useAppSelector((state) => state.posts);
  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(APIService.getPosts());
  }, [dispatch]);

  return (
    <Container className='mt-10'>
      {posts?.map((post,index)=>(
        <Fragment key={index}>
          <PostItem post={post}/>
          {index === posts.length - 1 ? null : <BorderLine/>}
        </Fragment>
    ))}
    </Container>
  );
}
