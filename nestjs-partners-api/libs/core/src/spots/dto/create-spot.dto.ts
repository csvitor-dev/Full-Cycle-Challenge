import { IsString, MaxLength } from 'class-validator';

export class CreateSpotDto {
  @IsString({
    message: 'name must be string',
  })
  @MaxLength(255, {
    message: 'name is too long',
  })
  name: string;
}
