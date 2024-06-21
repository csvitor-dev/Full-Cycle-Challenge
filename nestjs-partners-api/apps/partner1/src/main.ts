import { NestFactory } from '@nestjs/core';
import { Partner1Module } from './partner1.module';
import { ValidationExceptionFilter } from './validation-exception/validation-exception.filter';
import { CustomValidationPipe } from './custom-validation/custom-validation.pipe';

async function bootstrap() {
  const app = await NestFactory.create(Partner1Module);
  app.useGlobalPipes(new CustomValidationPipe());
  app.useGlobalFilters(new ValidationExceptionFilter());
  await app.listen(3000);
}
bootstrap();
