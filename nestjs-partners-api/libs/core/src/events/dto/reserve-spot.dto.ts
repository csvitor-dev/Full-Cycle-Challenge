import { TicketKind } from '@prisma/client';
import { Type } from 'class-transformer';
import { IsArray, IsEnum, ValidateNested } from 'class-validator';

export class ReserveSpotDto {
  @IsArray({
    message: 'type not recognized',
  })
  @ValidateNested({ each: true })
  @Type(() => String)
  spots: string[];

  @IsEnum(TicketKind, {
    message: 'type not recognized',
  })
  ticket_kind: TicketKind;

  email: string;
}
