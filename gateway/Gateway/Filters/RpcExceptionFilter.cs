using Grpc.Core;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Filters;

namespace Gateway.Filters;

public class RpcExceptionFilter : IExceptionFilter
{
    public void OnException(ExceptionContext context)
    {
        if (context.Exception is RpcException rpcException)
        {
            var (statusCode, title) = rpcException.Status.StatusCode switch
            {
                StatusCode.NotFound => (StatusCodes.Status404NotFound, "Resource Not Found"),
                _ => (StatusCodes.Status500InternalServerError, "Internal Server Error")
            };

            var response = new
            {
                error = title,
                detail = rpcException.Status.Detail,
                traceId = context.HttpContext.TraceIdentifier
            };

            context.Result = new ObjectResult(response)
            {
                StatusCode = statusCode
            };
            
            context.ExceptionHandled = true;
        }
    }
}