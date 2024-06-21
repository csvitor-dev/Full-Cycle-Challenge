import {
  IsISO8601,
  IsNumber,
  IsString,
  MaxLength,
  Min,
} from 'class-validator';

export class CreateEventDto {
  @IsString({
    message: 'name must be string',
  })
  @MaxLength(255, {
    message: 'name is too long',
  })
  name: string;

  @IsString({
    message: 'description must be string',
  })
  @MaxLength(255, {
    message: 'description is too long',
  })
  description: string;

  @IsISO8601()
  date: string;

  @IsNumber()
  @Min(0, {
    message: "negatives values aren't allowed",
  })
  price: number;
}
