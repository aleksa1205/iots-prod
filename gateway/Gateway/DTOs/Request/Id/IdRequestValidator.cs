using FluentValidation;

namespace Gateway.DTOs.Request.Id;

public class IdRequestValidator : AbstractValidator<IdRequest>
{
    public IdRequestValidator()
    {
        RuleFor(x => x.Id)
            .Must(id => Guid.TryParse(id, out _))
            .WithMessage("ID must be a valid GUID");
    }
}