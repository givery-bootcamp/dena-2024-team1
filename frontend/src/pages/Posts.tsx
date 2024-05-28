import { useEffect,Fragment } from 'react';
import { Link } from 'react-router-dom';

import { useAppDispatch, useAppSelector } from '../shared/hooks';
import { APIService } from '../shared/services';
import { Container } from '../shared/components/Container';

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
    <Link to={`/posts/${post.id}`} className="flex justify-between items-center">
      <div className="font-bold text-xl">
        {post.title}
      </div>
      <div>
        <div className='font-semi-bold text-gray-100 text-sm'>
          {post.userName}
        </div>
        <div className='text-gray-200 text-sm'>
          更新日時: {post.updatedAt}
        </div>
      </div>
    </Link>
    {index === posts.length - 1 ? null : <div className='flex justify-center py-3 items-center'><hr className="border-border w-full" /></div>}
    </Fragment>
    ))}
  </Container>
  );
}
