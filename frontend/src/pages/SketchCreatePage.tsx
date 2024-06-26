import { SketchCreateForm } from "~/features/sketches/SketchCreateForm";
import { Container } from "~/shared/components/Container";

export const SketchCreatePage = () => {
  return (
    <Container className="max-w-[900px] pt-10">
      <SketchCreateForm />
    </Container>
  );
};
