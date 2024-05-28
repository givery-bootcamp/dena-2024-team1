import { useEffect,Fragment } from "react";
import { Link } from "react-router-dom";

import { useAppDispatch, useAppSelector } from "../shared/hooks";
import { APIService } from "../shared/services";
import { Container } from "../shared/components/Container";

export function Posts() {
  const { posts } = useAppSelector((state) => state.posts);
  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(APIService.getPosts());
  }, [dispatch]);

  return (
  <Container className='mt-10'>
    {posts?.map((post,index)=>(
    <Fragment key={index}>
    <Link to={`/posts/${post.id}`} className="flex items-center justify-between">
      <div className="text-xl font-bold">
        {post.title}
      </div>
      <div>
        <div className='font-semi-bold text-sm text-gray-100'>
          {post.userName}
        </div>
        <div className='text-sm text-gray-200'>
          更新日時: {post.updatedAt}
        </div>
      </div>
    </Link>
    {index === posts.length - 1 ? null : <div className='flex items-center justify-center py-3'><hr className="w-full border-border" /></div>}
    </Fragment>
    ))}
  </Container>
  );
}
