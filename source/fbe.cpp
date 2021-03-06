/*!
    \file fbe.cpp
    \brief Fast binary encoding implementation
    \author Ivan Shynkarenka
    \date 11.04.2018
    \copyright MIT License
*/

#include "fbe.h"

namespace FBE {

void Import::AddImport(std::string* i)
{
    if (i == nullptr)
        yyerror("Import is null!");
    if (i->empty())
        yyerror("Import package is invalid!");

    // Check for duplicates
    auto it = std::find_if(imports.begin(), imports.end(), [i](auto item)->bool { return *item.get() == *i; });
    if (it != imports.end())
        yyerror("Duplicate import package " + *i);

    imports.push_back(std::shared_ptr<std::string>(i));
}

void EnumBody::AddValue(EnumValue* v)
{
    if (v == nullptr)
        yyerror("Enum is null!");
    if (v->name->empty())
        yyerror("Enum name is invalid!");

    // Check for duplicates
    auto it = std::find_if(values.begin(), values.end(), [v](auto item)->bool { return *item->name.get() == *v->name.get(); });
    if (it != values.end())
        yyerror("Duplicate enum name " + *v->name.get());

    values.push_back(std::shared_ptr<EnumValue>(v));
}

void FlagsBody::AddValue(FlagsValue* v)
{
    if (v == nullptr)
        yyerror("Flags is null!");
    if (v->name->empty())
        yyerror("Flags name is invalid!");
    if (!v->value || (!v->value->constant && !v->value->reference))
        yyerror("Flags value is invalid!");

    // Check for duplicates
    auto it = std::find_if(values.begin(), values.end(), [v](auto item)->bool { return *item->name.get() == *v->name.get(); });
    if (it != values.end())
        yyerror("Duplicate flags name " + *v->name.get());

    values.push_back(std::shared_ptr<FlagsValue>(v));
}

void StructBody::AddField(StructField* f)
{
    if (f == nullptr)
        yyerror("Struct field is null!");
    if (f->name->empty())
        yyerror("Struct field name is invalid!");
    if (f->type->empty())
        yyerror("Struct field type is invalid!");

    // Check for duplicates
    auto it = std::find_if(fields.begin(), fields.end(), [f](auto item)->bool { return *item->name.get() == *f->name.get(); });
    if (it != fields.end())
        yyerror("Duplicate struct field name " + *f->name.get());

    fields.push_back(std::shared_ptr<StructField>(f));
}

void Statements::AddStatement(Statement* st)
{
    if (st == nullptr)
        yyerror("Statement is null!");

    if (st->e)
        AddEnum(st->e);
    if (st->f)
        AddFlags(st->f);
    if (st->s)
        AddStruct(st->s);

    delete st;
}

void Statements::AddEnum(std::shared_ptr<EnumType>& e)
{
    if (e == nullptr)
        yyerror("Enum is null!");
    if (e->name->empty())
        yyerror("Enum name is invalid!");
    if (!e->body)
        yyerror("Enum is empty - " + *e->name.get());

    // Check for duplicates
    auto it = std::find_if(enums.begin(), enums.end(), [&e](auto item)->bool { return *item->name.get() == *e->name.get(); });
    if (it != enums.end())
        yyerror("Duplicate enum name " + *e->name.get());

    enums.push_back(e);
}

void Statements::AddFlags(std::shared_ptr<FlagsType>& f)
{
    if (f == nullptr)
        yyerror("Flags is null!");
    if (f->name->empty())
        yyerror("Flags name is invalid!");
    if (!f->body)
        yyerror("Flags is empty - " + *f->name.get());

    // Check for duplicates
    auto it = std::find_if(flags.begin(), flags.end(), [&f](auto item)->bool { return *item->name.get() == *f->name.get(); });
    if (it != flags.end())
        yyerror("Duplicate flags name " + *f->name.get());

    flags.push_back(f);
}

void Statements::AddStruct(std::shared_ptr<StructType>& s)
{
    if (s == nullptr)
        yyerror("Struct is null!");
    if (s->name->empty())
        yyerror("Struct name is invalid!");
    if (!s->body)
        yyerror("Struct is empty - " + *s->name.get());

    // Check for duplicates
    auto it = std::find_if(structs.begin(), structs.end(), [&s](auto item)->bool { return *item->name.get() == *s->name.get(); });
    if (it != structs.end())
        yyerror("Duplicate struct name " + *s->name.get());
    it = std::find_if(structs.begin(), structs.end(), [&s](auto item)->bool { return (item->type == s->type) && item->fixed && s->fixed; });
    if (it != structs.end())
        yyerror("Duplicate struct type " + std::to_string(s->type));

    structs.push_back(s);
}

int StructType::stype = 0;

StructType::StructType(int t, bool f) : type(t), fixed(f), keys(false)
{
    if (type < 0)
        yyerror("Struct type should not be negative!");

    if (!fixed)
    {
        if (type == 0)
            type = ++stype;
        else
            stype = type;
    }
}

std::shared_ptr<Package> Package::root = std::make_shared<Package>(0);

Package::Package(int o) : offset(o)
{
    if (offset < 0)
        yyerror("Package offset should not be negative!");
}

void Package::initialize()
{
    if (body)
    {
        for (const auto& child_s : body->structs)
        {
            // Add offset to all structs in the package
            if (!child_s->fixed)
                child_s->type += offset;

            // Find structs with key flag
            if (child_s->body)
            {
                for (const auto& field : child_s->body->fields)
                {
                    if (field->keys)
                    {
                        child_s->keys = true;
                        break;
                    }
                }
            }
        }
    }
}

} // namespace FBE
