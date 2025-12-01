using FluentValidation;

namespace Gateway.DTOs.Request.Time;

public class TimeRequestValidator : AbstractValidator<TimeRequest>
{
    public TimeRequestValidator()
    {
        RuleFor(x => x.StartTime)
            .Must(date => date.ToUnixTimeSeconds() > 0)
            .WithMessage("Start time must be greater than January 1st 1970");

        RuleFor(x => x.StartTime)
            .LessThan(x => x.EndTime)
            .WithMessage("Start time must be greater than End time");
    }
}