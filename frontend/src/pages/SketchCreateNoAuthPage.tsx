import { Link } from "react-router-dom";

import { Container } from "~/shared/components/Container";

export const SketchCreateNoAuthPage = () => {

  return (
    <Container className="pt-10">
      <Link to="/signin" className="underline">
        お絵描きするにはログインが必要です
      </Link>
    </Container>
  );
};
