import {
  ArgumentsHost,
  BadRequestException,
  Catch,
  ExceptionFilter,
  HttpStatus,
} from '@nestjs/common';

@Catch(BadRequestException)
export class ValidationExceptionFilter implements ExceptionFilter {
  catch(exception: BadRequestException, host: ArgumentsHost) {
    const context = host.switchToHttp();
    const response = context.getResponse();
    const status = HttpStatus.UNPROCESSABLE_ENTITY;

    response.status(status).json({
      statusCode: status,
      message: 'problems found',
      errors: exception.getResponse()['message'],
    });
  }
}
